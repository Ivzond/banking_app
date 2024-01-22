import React, {FC, useState} from 'react';
import { Modal, Box, IconButton } from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';
import LoginForm from '../forms/LoginForm.tsx'; // Ваш компонент формы логина
import { useTheme } from '@mui/material/styles';

interface AuthModalProps{
    isOpen: boolean,
    onClose: any,
}
const AuthModal : FC<AuthModalProps> = ({ isOpen, onClose }) => {
    const theme = useTheme();

    return (
        <Modal
            open={isOpen}
            onClose={onClose}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Box
                sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    border: '2px solid #000',
                    boxShadow: 24,
                    p: 4,
                    borderRadius: 4,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <IconButton
                    aria-label="close"
                    onClick={onClose}
                    sx={{ position: 'absolute', top: theme.spacing(1), right: theme.spacing(1) }}
                >
                    <CloseIcon />
                </IconButton>
                <LoginForm />
            </Box>
        </Modal>
    );
};

export default AuthModal;
