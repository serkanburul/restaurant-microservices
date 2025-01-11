import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Calendar, Clock, Users, ArrowLeft, Trash2 } from "lucide-react";
import dayjs from 'dayjs';
import 'dayjs/locale/en';
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
} from "@/components/ui/alert-dialog";

const ReservationDetail = () => {
    const { token } = useParams();
    const navigate = useNavigate();

    const [reservation, setReservation] = useState(null);
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchReservation = async () => {
            setIsLoading(true);
            try {
                const response = await fetch(`http://localhost:8080/reservation/${token}`);
                if (!response.ok) throw new Error('Reservation not found');

                const data = await response.json();
                setReservation(data);
            } catch (err) {
                setError(err.message);
            } finally {
                setIsLoading(false);
            }
        };

        fetchReservation();
    }, [token]);

    const handleDelete = async () => {
        try {
            const response = await fetch(`http://localhost:8080/reservation/${token}`, {
                method: 'DELETE',
            });

            if (response.ok) {
                setIsDeleteDialogOpen(false);
                navigate('/reservation');
            }
        } catch (error) {
            console.error('Error deleting reservation:', error);
        }
    };

    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen">
                <div className="text-center">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="flex flex-col items-center justify-center min-h-screen">
                <div className="text-red-500 mb-4">{error}</div>
                <Button variant="outline" onClick={() => navigate('/reservation')}>
                    <ArrowLeft className="mr-2 h-4 w-4" />
                    Back
                </Button>
            </div>
        );
    }

    return (
        <div className="container mx-auto p-4 max-w-4xl mt-24">
            <Button
                variant="outline"
                onClick={() => navigate('/reservation')}
                className="mb-6"
            >
                <ArrowLeft className="mr-2 h-4 w-4" />
                Go Back To Reservations
            </Button>

            <Card>
                <CardHeader className="flex flex-row items-center justify-between">
                    <CardTitle>Reservation Detail</CardTitle>
                    <div className="flex space-x-2">
                        <Button
                            variant="destructive"
                            size="icon"
                            onClick={() => setIsDeleteDialogOpen(true)}
                        >
                            <Trash2 className="h-4 w-4" />
                        </Button>
                    </div>
                </CardHeader>

                <CardContent>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                        <div className="space-y-6">
                            <div className="flex items-center space-x-3">
                                <Users className="h-5 w-5 text-gray-500" />
                                <div>
                                    <div className="text-sm text-gray-500">Reservation Owner</div>
                                    <div className="font-medium">{reservation.name}</div>
                                </div>
                            </div>

                            <div className="flex items-center space-x-3">
                                <Calendar className="h-5 w-5 text-gray-500" />
                                <div>
                                    <div className="text-sm text-gray-500">Date</div>
                                    <div className="font-medium">
                                        {dayjs(reservation["reservation_date"]).format('DD MMMM YYYY')}
                                    </div>
                                </div>
                            </div>

                            <div className="flex items-center space-x-3">
                                <Clock className="h-5 w-5 text-gray-500" />
                                <div>
                                    <div className="text-sm text-gray-500">Time</div>
                                    <div className="font-medium">{reservation["time_slot"]}</div>
                                </div>
                            </div>

                            <div className="flex items-center space-x-3">
                                <Users className="h-5 w-5 text-gray-500" />
                                <div>
                                    <div className="text-sm text-gray-500">Guests</div>
                                    <div className="font-medium">{reservation["guests"]} People</div>
                                </div>
                            </div>
                        </div>

                        <Card className="bg-gray-50">
                            <CardHeader>
                                <CardTitle className="text-base">Reservation Information</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-2">
                                <div className="text-sm">
                                    <span className="text-gray-500">Reservation No:</span>
                                    <span className="ml-2 font-mono">{token}</span>
                                </div>
                                <div className="text-sm">
                                    <span className="text-gray-500">Created At:</span>
                                    <span className="ml-2">
                                        {dayjs(reservation["created_at"]).format('DD.MM.YYYY HH:mm')}
                                    </span>
                                </div>
                            </CardContent>
                        </Card>
                    </div>
                </CardContent>
            </Card>

            <AlertDialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
                <AlertDialogContent>
                    <AlertDialogHeader>
                        <AlertDialogTitle>Delete Reservation</AlertDialogTitle>
                        <AlertDialogDescription>
                            This action cannot be undone. Are you sure you want to delete this reservation?
                        </AlertDialogDescription>
                    </AlertDialogHeader>
                    <AlertDialogFooter>
                        <AlertDialogCancel>Cancel</AlertDialogCancel>
                        <AlertDialogAction onClick={handleDelete} className="bg-red-500">
                            Delete
                        </AlertDialogAction>
                    </AlertDialogFooter>
                </AlertDialogContent>
            </AlertDialog>
        </div>
    );
};

export default ReservationDetail;