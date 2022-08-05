import { Routes, Route } from "react-router-dom";
import { Campaigns } from "../../pages/Campaigns";
import { Create } from "../../pages/Campaigns/Create";
import { Edit } from "../../pages/Campaigns/Edit";
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
            <Route path="dashboard" element={<Private><Dashboard /></Private>} />
            <Route path="campaigns" >
                <Route path="create" element={<Create />} />
                <Route path="edit/:campaignId" element={<Edit />} />
                <Route index element={<Campaigns />} />
            </Route>
            <Route path="*" element={<NotFound />} />
        </Routes>
    );
}
