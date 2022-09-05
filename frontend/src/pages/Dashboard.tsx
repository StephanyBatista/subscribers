import { Layout } from "../components/templates/Layout";
import { useAuth } from "../hooks/useAuth"

export function Dashboard() {
    const { user } = useAuth();

    return (
        <Layout>
            <h1>DashBoard</h1>
        </Layout>
    )
}