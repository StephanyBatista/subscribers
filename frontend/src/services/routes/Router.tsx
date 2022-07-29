import { Routes, Route } from "react-router-dom";
import { Campaigns } from "../../pages/Campaigns";
import { Create } from "../../pages/Campaigns/Create";
import { Dashboard } from "../../pages/Dashboard";
import { Register } from "../../pages/Register";
import { SingIn } from "../../pages/SingIn";
import { NotFound } from "./NotFound";
import { Private } from "./Private";

export function AppRoute() {


    return (
        <Routes>
            <Route path="/" element={<SingIn />} />
            <Route path="register" element={<Register />} />
            <Route path="singin" element={<SingIn />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="campaigns" >
                <Route path="create" element={<Create />} />
                <Route index element={<Campaigns />} />
            </Route>
            <Route path="*" element={<NotFound />} />
        </Routes>
    );
}