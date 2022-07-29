import { Routes, Route } from "react-router-dom";
import { Dashboard } from "../../pages/Dashboard";
import { Register } from "../../pages/Register";
import { SingIn } from "../../pages/SingIn";
import { NotFound } from "./NotFound";
import { Private } from "./Private";

export function AppRoute() {


    return (
        <Routes>
            <Route path="/" element={<SingIn />} />
            <Route path="/register" element={<Register />} />
            <Route path="/dashboard" element={<Private><Dashboard /></Private>} />
            <Route path="/singin" element={<SingIn />} />
            <Route path="*" element={<NotFound />} />
        </Routes>
    );
}
