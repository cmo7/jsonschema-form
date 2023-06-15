import { useParams } from "react-router-dom"
import UserLayout from "../layouts/user/user-layout";

export default function UserProfile() {
    const {userId} = useParams<"userId">();
    return (
        <UserLayout>
            <h1>Perfil de usuario {userId}</h1>
        </UserLayout>
    )
}