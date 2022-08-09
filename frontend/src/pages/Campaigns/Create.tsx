import { Link, Box, Button, Flex, Textarea, FormLabel, Heading, IconButton, Stack, Switch, Icon, useToast, FormErrorMessage, Alert, AlertIcon, AlertDescription, FormControl } from "@chakra-ui/react";
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { BiArrowBack } from "react-icons/bi";
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";

interface FormProps {
    name: string;
    from: string;
    body: string;
    subject: string;
}

const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    from: Yup.string().email().required('E-mail de origem é obrigatório'),
    body: Yup.string().required('A texto é obrigatório'),
    subject: Yup.string().required('O assunto é obrigatório'),
    //active: Yup.boolean()
    //.required("The terms and conditions must be accepted.s")
    //  .oneOf([true], "The terms and conditions must be accepted.")
})
export function Create() {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();

    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        console.log(values)
        api.post('/campaigns', {
            name: values.name,
            from: values.from,
            body: values.body,
            subject: values.subject
        }).then((response) => {
            if (response.status === 201) {
                toast({
                    description: "Campanha cadastrada com sucesso!",
                    status: 'success',
                    duration: 5000,
                    isClosable: true
                });
                navigate('/campaigns');
            }
        }).catch((err) => {
            console.log(err)
        });

    }


    return (
        <Layout>
            <Flex
                bg="gray.800"
                maxH={700}
                p="8"
                w="100vw"
                maxW={600}
                m="0 auto"
                flexDirection="column"
                borderRadius={8}>
                <Flex
                    onSubmit={handleSubmit(onHandleSubmit)}
                    as="form"
                    flexDirection="column"
                    justify="space-between"
                >
                    <Stack spacing={20}>
                        <Flex justify="space-between" align="center">
                            <Heading fontSize="xl">Criar uma campanha</Heading>
                            <Link
                                as={ReactLink}
                                to="/campaigns"
                            >
                                <Icon as={BiArrowBack} fontSize="2xl" />
                            </Link>
                        </Flex>
                        <Stack>
                            <Input
                                {...register('name')}
                                type="text"
                                label="Nome"
                                error={errors.name}
                            />
                            <Input
                                {...register('from')}
                                type="email"
                                label="De"
                                placeholder="exemplo@dominio.com"
                                error={errors.from}
                            />
                            <Input
                                {...register('subject')}
                                type="text"
                                label="Assunto"
                                error={errors.subject}
                            />
                            <FormControl>
                                <FormLabel>Texto</FormLabel>
                                <Textarea
                                    {...register('body')}
                                    resize="none"
                                    bg="gray.950"
                                    border="none"
                                />
                                {errors.body && (
                                    <Alert bg="transparent" color="red.600" fontSize="0.875rem">
                                        <AlertDescription>
                                            {errors.body?.message}
                                        </AlertDescription>
                                    </Alert>

                                )}
                            </FormControl>
                        </Stack>
                    </Stack>
                    <Flex mt="10">
                        <Link
                            py="2"
                            px="4"
                            borderRadius="6"
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="gray.600"
                            as={ReactLink}
                            mr={3}
                            to="/campaigns"
                        >
                            Voltar
                        </Link>
                        <Button
                            type="submit"
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="blue.900"
                        >Salvar
                        </Button>
                    </Flex>
                </Flex>


            </Flex>
        </Layout>
    )
}