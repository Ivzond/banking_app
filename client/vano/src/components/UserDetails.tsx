// src/components/UserDetails.tsx
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const UserDetails: React.FC = () => {
    const [userId, setUserId] = useState<string>('');
    const [userData, setUserData] = useState<any>(null);

    const fetchUserDetails = async () => {
        try {
            const response = await axios.get(`http://localhost:8080/user/${userId}`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('jwt')}`,
                },
            });

            setUserData(response.data);
        } catch (error: any) {
            console.log(error.message);
        }
    };

    useEffect(() => {
        if (userId) {
            fetchUserDetails();
        }
    }, [userId]);

    return (
        <div>
            <h2>User Details</h2>
            <input
                type="text"
                placeholder="User ID"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
            />
            {userData && (
                <div>
                    <p>ID: {userData.ID}</p>
                    <p>Username: {userData.Username}</p>
                    <p>Email: {userData.Email}</p>
                </div>
            )}
        </div>
    );
};

export default UserDetails;
