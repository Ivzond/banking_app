import React, { FC, useState } from 'react';
import { AppBar, List, ListItem, ListItemText } from '@mui/material';
import HomeIcon from '@mui/icons-material/Home';
import AuthModal from '../auth/AuthModal'; // Assuming you have an AuthModal for login
import RegisterModal from '../auth/Register'; // Import the RegisterModal
import { Link } from 'react-router-dom';

function NavPanel() {
    const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
    const [isRegisterModalOpen, setIsRegisterModalOpen] = useState(false);

    const openLoginModal = () => {
        setIsLoginModalOpen(true);
    };

    const closeLoginModal = () => {
        setIsLoginModalOpen(false);
    };

    const openRegisterModal = () => {
        setIsRegisterModalOpen(true);
    };

    const closeRegisterModal = () => {
        setIsRegisterModalOpen(false);
    };

    return (
        <AppBar className="nav-panel">
            <List className="nav-panel-list">
                <Link to={'/'} className="link-style">
                    <ListItem className="nav-panel-item nav-panel-item-left">
                        <HomeIcon />
                        <ListItemText disableTypography={true} className="nav-panel-item-text-icon">
                            Home
                        </ListItemText>
                    </ListItem>
                </Link>
                <ListItem>
                    <button onClick={openLoginModal}>Login</button>
                    <AuthModal isOpen={isLoginModalOpen} onClose={closeLoginModal} />
                </ListItem>
                <ListItem>
                    <button onClick={openRegisterModal}>Register</button>
                    <RegisterModal isOpen={isRegisterModalOpen} onClose={closeRegisterModal} />
                </ListItem>
            </List>
        </AppBar>
    );
}

export default NavPanel;
