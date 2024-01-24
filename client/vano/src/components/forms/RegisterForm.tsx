import React from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import { InputLabel, TextField } from '@mui/material';
import { IUserRegister } from '../../interfaces/IUser';
import { userRegister } from '../../utils/userAuth';

const RegisterForm: React.FC = () => {
    const { register, handleSubmit } = useForm<IUserRegister>();

    const onSubmit: SubmitHandler<IUserRegister> = async (data: IUserRegister) => {
        try {
            await userRegister(data);
        } catch (error: any) {
            console.log(error.message);
        }
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="register-form">
            <InputLabel className="register-form-username-label">Username</InputLabel>
            <TextField
                className="register-form-username-input"
                size="small"
                id="username"
                type="text"
                placeholder="Your username"
                {...register('username', { required: true })}
            />
            <InputLabel className="register-form-email-label">Email </InputLabel>
            <TextField
                className="register-form-email-input"
                size="small"
                id="email"
                type="email"
                placeholder="your_email@test.com"
                {...register('email', { required: true })}
            />
            <InputLabel className="register-form-password-label">Password</InputLabel>
            <TextField
                className="register-form-password-input"
                size="small"
                id="password"
                placeholder="Your password"
                type="password"
                {...register('password', { required: true })}
            />

            <button type="submit" className="register-form-submit-button">
                Register
            </button>
        </form>
    );
};

export default RegisterForm;
