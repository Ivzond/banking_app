import React, {FC, useState} from "react";
import { Link } from 'react-router-dom';
import Login from "../auth/Login.tsx";

import Register from "../auth/Register.tsx";
import { AppBar, List, ListItem, ListItemText } from "@mui/material";
import HomeIcon from '@mui/icons-material/Home';


import AuthModal from "../auth/Login.tsx";

function NavPanel() {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const openModal = () => {
        setIsModalOpen(true);
    };

    const closeModal = () => {
        setIsModalOpen(false);
    };
    return (
        <AppBar className='nav-panel'>
            <List className='nav-panel-list'>
                <Link to={'/'} className='link-style'>
                    <ListItem className='nav-panel-item nav-panel-item-left'>
                        <HomeIcon
                        />
                       <ListItemText disableTypography={true}  className='nav-panel-item-text-icon' >Home</ListItemText>)

                    </ListItem>
                </Link>
                <ListItem>
                    <button onClick={openModal} >Login</button>

                    <AuthModal isOpen={isModalOpen} onClose={closeModal}/>
                </ListItem>

            </List>
        </AppBar>
    );
}
export default NavPanel;