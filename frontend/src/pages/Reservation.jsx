import React, { useState, useEffect } from 'react';
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Calendar } from "@/components/ui/calendar";
import dayjs from 'dayjs';
import 'dayjs/locale/tr';
import { Clock } from "lucide-react";

dayjs.locale('tr');

const CreateReservation = () => {
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        table_no: '',
        reservation_date: '',
        time_slot: '',
    });

    const [selectedDate, setSelectedDate] = useState(null);
    const [bookedTimeslots, setBookedTimeslots] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [bookingsByTimeSlot, setBookingsByTimeSlot] = useState({});
    const [guests, setGuests] = useState('');
    const [slots, setSlots] = useState([]);

    const hasTenOccurences = (array, timeSlot) => {
        const count = array.filter(item => item === timeSlot).length;
        return count >= 10;
    };

    const getTableNo = (guestCount, timeSlot) => {
        const guests = Number(guestCount);
        if (!Number.isInteger(guests) || guests <= 0 || guests > 5) {
            throw new Error(`Invalid guest count: ${guests}`);
        }

        const tablesForTimeSlot = bookingsByTimeSlot[timeSlot] || [];

        const start = (guests - 1) * 10 + 1;
        const end = guests * 10;

        for (let i = start; i <= end; i++) {
            if (!tablesForTimeSlot.includes(i)) {
                return i;
            }
        }

        throw new Error("No available tables for this time slot");
    };

    const resetBookingStates = () => {
        setBookedTimeslots([]);
        setBookingsByTimeSlot({});
        setFormData(prev => ({ ...prev, time_slot: '' }));
    };

    const handleDateChange = (date) => {
        setSelectedDate(date);
        resetBookingStates();
    };

    const handleGuestsChange = (e) => {
        setGuests(e.target.value);
        resetBookingStates();
    };

    useEffect(() => {
        const fetchSlots = async () => {
            try {
                const response = await fetch('http://localhost:8080/reservation/slots');
                if (response.ok) {
                    const data = await response.json();
                    setSlots(data);
                } else {
                    console.error('Failed to fetch slots');
                }
            } catch (error) {
                console.error('Error fetching slots:', error);
            }
        };

        fetchSlots();
    }, []);

    useEffect(() => {
        const fetchBookedTimeslots = async () => {
            if (!selectedDate || !guests || guests < 1) return;

            setIsLoading(true);
            try {
                const formattedDate = dayjs(selectedDate).format('YYYY-MM-DD');
                const response = await fetch(`http://localhost:8080/reservation/${formattedDate}/${guests}`);

                if (response.ok) {
                    const reservations = await response.json();
                    console.log('Received reservations:', reservations);

                    if (reservations) {
                        const bookingsByTime = {};
                        const bookedSlots = [];

                        reservations.forEach(reservation => {
                            const timeSlot = reservation["StartTime"];
                            bookedSlots.push(timeSlot);

                            if (!bookingsByTime[timeSlot]) {
                                bookingsByTime[timeSlot] = [];
                            }
                            bookingsByTime[timeSlot].push(reservation["TableNo"]);
                        });

                        setBookedTimeslots(bookedSlots);
                        setBookingsByTimeSlot(bookingsByTime);

                        if (hasTenOccurences(bookedSlots, formData.time_slot)) {
                            console.log(`Current time slot ${formData.time_slot} is fully booked, clearing selection`);
                            setFormData(prev => ({ ...prev, time_slot: '' }));
                        }
                    }
                }
            } catch (error) {
                console.error('Reservation Control Error: ', error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchBookedTimeslots();
    }, [selectedDate, guests]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const assignedTable = getTableNo(Number(guests), formData.time_slot);
            const response = await fetch('http://localhost:8080/reservation', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    ...formData,
                    table_no: assignedTable,
                    reservation_date: dayjs(selectedDate).format('YYYY-MM-DD')
                }),
            });

            if (response.ok) {
                alert("Reservation Created Successfully!");
                setFormData({ name: '', email: '', table_no: '', reservation_date: '', time_slot: '' });
                setGuests('');
                setSelectedDate(null);
                setBookedTimeslots([]);
                setBookingsByTimeSlot({});
            } else {
                throw new Error('Error while creating bookedTimeslot');
            }
        } catch (error) {
            console.error('Reservation Creation Error:', error);
            alert(`Reservation Could Not Be Created: ${error.message}`);
        }
    };

    return (
        <Card className="w-full max-w-4xl mx-auto mt-24 mb-4">
            <CardHeader>
                <CardTitle>New Reservation</CardTitle>
            </CardHeader>
            <CardContent>
                <form onSubmit={handleSubmit} className="space-y-6">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <div className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="name">Name</Label>
                                <Input
                                    id="name"
                                    value={formData.name}
                                    onChange={(e) => setFormData({...formData, name: e.target.value})}
                                    required
                                    placeholder="Enter your name"
                                />
                            </div>

                            <div className="space-y-2">
                                <Label htmlFor="email">Mail</Label>
                                <Input
                                    id="email"
                                    type="email"
                                    value={formData.email}
                                    onChange={(e) => setFormData({...formData, email: e.target.value})}
                                    required
                                    placeholder="Enter your mail address"
                                />
                            </div>

                            <div className="space-y-2">
                                <Label htmlFor="guests">People</Label>
                                <Input
                                    id="guests"
                                    type="number"
                                    value={guests}
                                    onChange={handleGuestsChange}
                                    required
                                    min="1"
                                    max="5"
                                    placeholder="Please enter a guest"
                                />
                            </div>

                            <div className="space-y-2">
                                <Label>Date</Label>
                                <div className="border rounded-md p-4">
                                    <Calendar
                                        mode="single"
                                        selected={selectedDate}
                                        onSelect={handleDateChange}
                                        disabled={{before: new Date()}}
                                        className="w-full"
                                    />
                                </div>
                            </div>
                        </div>

                        {/* Sağ Kolon - Saat Seçimi */}
                        <div className="space-y-4">
                            <Label>Time Selection</Label>
                            {(!selectedDate || !guests) ? (
                                <div className="text-center p-8 border rounded-md bg-gray-50">
                                    <Clock className="mx-auto h-12 w-12 text-gray-400"/>
                                    <p className="mt-2 text-sm text-gray-500">
                                        Please select a date and number of people to see available times.
                                    </p>
                                </div>
                            ) : isLoading ? (
                                <div className="text-center p-8 border rounded-md">
                                    <p>Checking available times...</p>
                                </div>
                            ) : (
                                <RadioGroup
                                    value={formData.time_slot}
                                    onValueChange={(value) => setFormData({...formData, time_slot: value})}
                                    className="grid grid-cols-3 gap-2"
                                >
                                    {slots.map((time) => {
                                        const isDisabled = hasTenOccurences(bookedTimeslots, time);
                                        return (
                                            <div key={time}>
                                                <RadioGroupItem
                                                    value={time}
                                                    id={time}
                                                    disabled={isDisabled}
                                                    className="peer sr-only"
                                                />
                                                <Label
                                                    htmlFor={time}
                                                    className={`flex items-center justify-center px-4 py-2 border rounded-md text-sm
                                                        ${isDisabled
                                                        ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
                                                        : 'peer-checked:bg-primary peer-checked:text-white hover:bg-gray-50 cursor-pointer'
                                                    }`}
                                                >
                                                    {time.slice(0, 5)}
                                                    {isDisabled && <span className="ml-1 text-xs">(Full)</span>}
                                                </Label>
                                            </div>
                                        );
                                    })}
                                </RadioGroup>
                            )}
                        </div>
                    </div>

                    <Button
                        type="submit"
                        className="w-full"
                        disabled={!formData.name || !formData.email || !guests || !selectedDate || !formData.time_slot}
                    >
                        Create Reservation
                    </Button>
                </form>
            </CardContent>
        </Card>
    );
};

export default CreateReservation;
