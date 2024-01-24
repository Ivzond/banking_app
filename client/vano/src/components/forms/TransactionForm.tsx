// src/components/forms/TransactionForm.tsx
import React, { useState } from 'react';
import axios from 'axios';

const TransactionForm: React.FC = () => {
    const [userId, setUserId] = useState<string>('');
    const [from, setFrom] = useState<number>(0);
    const [to, setTo] = useState<number>(0);
    const [amount, setAmount] = useState<number>(0);

    const handleTransaction = async () => {
        try {
            const data = {
                UserID: parseInt(userId),
                From: from,
                To: to,
                Amount: amount,
            };

            const response = await axios.post('http://localhost:8080/transaction', data, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('jwt')}`,
                },
            });

            // Handle the response as needed
        } catch (error: any) {
            console.log(error.message);
        }
    };

    return (
        <div>
            <h2>Transaction</h2>
            <input
                type="text"
                placeholder="User ID"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
            />
            <input
                type="number"
                placeholder="From"
                value={from}
                onChange={(e) => setFrom(Number(e.target.value))}
            />
            <input
                type="number"
                placeholder="To"
                value={to}
                onChange={(e) => setTo(Number(e.target.value))}
            />
            <input
                type="number"
                placeholder="Amount"
                value={amount}
                onChange={(e) => setAmount(Number(e.target.value))}
            />
            <button onClick={handleTransaction}>Make Transaction</button>
        </div>
    );
};

export default TransactionForm;
