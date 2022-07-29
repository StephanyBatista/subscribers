import { useAuth } from "../hooks/useAuth"

export function Dashboard() {
    const { user } = useAuth();
    console.log(user);
    return (
        <h1>Dashboard</h1>
    )
}