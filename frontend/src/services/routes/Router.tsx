import { Routes, Route } from "react-router-dom";
import { Campaigns } from "../../pages/Campaigns";
import { Create } from "../../pages/Campaigns/Create";
import { Edit } from "../../pages/Campaigns/Edit";
import { Contacts } from "../../pages/Contacts";
import { CreateContact } from "../../pages/Contacts/Create";
import { EditContact } from "../../pages/Contacts/Edit";
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
            <Route path="contacts" >
                <Route path="create" element={<CreateContact />} />
                <Route path="edit/:contactId" element={<EditContact />} />
                <Route index element={<Contacts />} />
            </Route>
            <Route path="*" element={<NotFound />} />
        </Routes>
    );
}
