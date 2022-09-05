import { Button, Flex, Heading, HStack, Link, Stack, useToast } from "@chakra-ui/react";
import axios from "axios";
import { Link as ReactLink, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom"
import { Layout } from "../../components/templates/Layout";
import { api } from "../../services/apiClient";

export function CancelSubscription() {
    const [contact, setContact] = useState();
    const [isFetching, setIsFetching] = useState(true);
    const [isLoading, setIsLoading] = useState(false);
    const { contactId } = useParams();

    const toast = useToast();
    const navigate = useNavigate();
    const handleCancelSubscription = async () => {
        api.patch(`contacts/${contactId}/cancel`)
            .then((response) => {
                toast({
                    description: 'Inscrição cancelada.',
                    status: 'success'
                });
                navigate('/singin');
            }).catch(err => console.log(err))
            .finally(() => setIsLoading(false));
    }
    // useEffect(() => {
    //     const controller = new AbortController();
    //     axios.get(`${import.meta.env.VITE_API_BACKEND}/contacts/${contactId}`, { signal: controller.signal })
    //         .then((response) => {
    //             console.log(response);
    //         }).catch((err) => console.log(err.response))
    //         .finally(() => setIsFetching(false));
    //     return () => { controller.abort() };
    // })
    return (
        <Layout>
            <Flex
                w="100vw"
                maxW={1480}
                bg="gray.800"
                px="8"
                py="8"
                flex="1"
                justify="center"
                align="center"
                borderRadius={8}
                flexDirection="column"
            >
                <Stack spacing="20"  >
                    <Stack spacing="8">
                        <Heading>Cancelar inscrição</Heading>
                        <Heading fontSize="xl">Deseja cancelar sua inscrição na plataforma ?</Heading>
                    </Stack>
                    <HStack spacing={8}>
                        <Button
                            _hover={{ color: 'white', bg: 'red.900' }}
                            onClick={handleCancelSubscription}
                            variant="unstyled"
                            width="100%"
                            display="block"
                            fontSize="lg"
                            fontWeight="bold"
                            isLoading={isLoading}
                        >Desejo cancelar</Button>
                        <Link
                            width="100%"
                            as={ReactLink}
                            to="/singin"
                        >Entrar na minha conta</Link>
                    </HStack>
                </Stack>
            </Flex>
        </Layout>
    )
}