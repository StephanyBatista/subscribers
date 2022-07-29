import { Box, Flex, Heading, Table, Th, Thead, Tr } from "@chakra-ui/react";
import { Layout } from "../../components/templates/Layout";

export function Campaigns() {
    return (
        <Layout>
            <Flex flexDirection="column">
                <Heading fontSize="2xl">Campanhas</Heading>
                <Box overflowY="auto">
                    <Table colorScheme='whiteAlpha'>
                        <Thead>
                            <Tr>
                                <Th></Th>
                            </Tr>
                        </Thead>
                    </Table>
                </Box>
            </Flex>
        </Layout>
    )
}