import {
    Box,
    useColorModeValue
} from '@chakra-ui/react';
import UserLayout from '../layouts/user/user-layout';

export default function Root() {
    return (
        <UserLayout>
            <Box>
                <h1>Soy Root</h1>
            </Box>
        </UserLayout>
    )
}