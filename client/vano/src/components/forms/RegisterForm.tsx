import React, {useState,useEffect} from 'react';
import { useForm,SubmitHandler, FieldValues } from 'react-hook-form';
import {InputLabel, TextField} from "@mui/material";
import {IUserRegister} from "../../interfaces/IUser";
import {userRegister} from "../../utils/userAuth.tsx";


const RegisterForm: React.FC = () => {
    const {register, handleSubmit } = useForm<IUserRegister>();

    const onSubmit: SubmitHandler<IUserRegister> = async (data:IUserRegister) => {
        try{
            await userRegister(data);
        }
        catch(error:any){
            console.log(error.message);
        }
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className='login-form'>
            <InputLabel  className='login-form-username-label'>Username</InputLabel>
            <TextField
                className='login-form-username-input'
                size='small'
                id="username"
                type='text'
                placeholder="vasya"
                {...register('username', { required: true })}
            />
            <InputLabel  className='login-form-Email-label'>Email </InputLabel>
            <TextField
                className='login-form-email-input'
                size='small'
                id="email"
                type='email'
                placeholder="vasya.pupkin@mail.ru"
                {...register('email', { required: true })}
            />
            <InputLabel  className='login-form-password-label'>Password</InputLabel>
            <TextField
                className='login-form-password-input'
                size='small'
                id='password'
                placeholder=""
                type='password'
                {...register('password', { required: true })}
            />


            <button type="submit" className='login-form-submit-button'>Register</button>
        </form>
    );
};

export default RegisterForm;