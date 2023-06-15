import { useEffect, useState } from "react";

interface UserListElement {
    id: number,
    created_at: string,
    updated_at: string,
    email: string,
    friends: UserListElement[],
}

export default function UserList() {
    const [userList, setUserList] = useState<UserListElement[]>([]);

    useEffect(() => {
        (async () => {
            const response = await fetch("http://localhost:8080/api/user/all");
            const jsonData = await response.json();
            setUserList(jsonData)
        })()
    }, [])

    return (
        <>
        {userList.map(u => <UserBadge email={u.email} friends={u.friends} />)}
        </>
    )
}

interface UserBadgeProps {
    email: string,
    friends: UserBadgeProps[],
}

function UserBadge(props: UserBadgeProps) {
    const {
        email,
        friends
    } = props;
    return (
        <>
            <h4>{email}</h4>
            {friends.map(f => <UserBadge email={f.email} friends={f.friends} />)}
        </>
    )
}