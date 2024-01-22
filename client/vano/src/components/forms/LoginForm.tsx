import React, {useState,useEffect} from 'react';
import { useForm,SubmitHandler, FieldValues } from 'react-hook-form';
import {InputLabel, TextField} from "@mui/material";
import {IUserLogin} from "../../interfaces/IUser";
import {userLogin} from "../../utils/userAuth.tsx";


const LoginForm: React.FC = () => {
    const {register, handleSubmit } = useForm<IUserLogin>();

    const onSubmit: SubmitHandler<IUserLogin> = async (data:IUserLogin) => {
        try{
            await userLogin(data);
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
            <InputLabel  className='login-form-password-label'>Password</InputLabel>
            <TextField
                className='login-form-password-input'
                size='small'
                id='password'
                placeholder=""
                type='password'
                {...register('password', { required: true })}
            />


            <button type="submit" className='login-form-submit-button'>Login</button>
        </form>
    );
};

export default LoginForm;