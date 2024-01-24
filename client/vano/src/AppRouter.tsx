import {BrowserRouter, Routes, Route} from 'react-router-dom'
import MainPage from "./components/pages/MainPage.tsx";
import NavPanel from "./components/common/NavPanel.tsx";

const AppRouter = () => {
    return(
        <BrowserRouter>
            <NavPanel/>
            <Routes>
                <Route path="/" element={<MainPage/>}/>
            </Routes>
        </BrowserRouter>
    )
}
export default AppRouter;