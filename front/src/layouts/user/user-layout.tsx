import { PropsWithChildren } from "react"
import Header from "./header"
import Footer from "./footer"
import { Box, Flex, VStack} from "@chakra-ui/react"
import Sidebar from "./sidebar"

export default function UserLayout(props: PropsWithChildren) {
    return (
        <>
            <Flex>
                <Sidebar />
                <Box flex='1'>
                    <VStack align="stretch">
                    <Header />
                    <Box>
                        {props.children}
                    </Box>
                    <Footer />
                    </VStack>
                </Box>
            </Flex>
        </>
    )
}