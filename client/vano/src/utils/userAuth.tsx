import React, {useState,useEffect} from 'react';
import {IUserLogin,IUserRegister} from "../interfaces/IUser.ts";
import axios from 'axios';
export async function userLogin (data:IUserLogin) {
    try {
        const response = await axios.post('http://localhost:8080/login', data);
        const responseData = response.data;
        localStorage.setItem('jwt',responseData.jwt);
        localStorage.setItem('user',JSON.stringify(responseData.data));

    }
    catch (error: any)
    {
        alert("Error while login. Try again");

    }
};
export async function userRegister (data:IUserRegister) {
    try {
        const response = await axios.post('http://localhost:8080/register', data);
        alert("Register completed");

    }
    catch (error: any)
    {
        alert("Error while registered. Try again");
        
    }
};
