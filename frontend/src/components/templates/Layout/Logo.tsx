import { Box, Heading, Link } from "@chakra-ui/react";
import { Link as ReactLink } from "react-router-dom";

export function Logo() {
    return (
        <Link
            _hover={{ textDecoration: 'none' }}
            as={ReactLink}
            to="/dashboard"
            ml="4"
        >
            <Box >
                <Heading fontSize="xl" as="span" >Subs</Heading>
                <Heading fontSize="xl" color="blue.900" as="span" >criber</Heading>
            </Box>
        </Link>
    )
}