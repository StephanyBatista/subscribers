import { Modal, ModalContent, ModalHeader, ModalCloseButton, ModalFooter, Button, Stack, useToast, Flex } from '@chakra-ui/react';
import { Input } from '../utils/Input';
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { useNavigate } from 'react-router-dom';
import { api } from '../../services/apiClient';

interface FormProps {
    name: string;
    email: string;
}


interface ClientModalProps {
    isOpen: boolean;
    campaignId: string | undefined;
    onClose: () => void;
    onUpdateState: () => void;
}

const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    email: Yup.string().email().required('E-mail é obrigatório'),
});



export function ClientModal({ isOpen, onClose, campaignId, onUpdateState }: ClientModalProps) {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();

    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        console.log(values)
        const response = await api.post('clients', {
            name: values.name,
            email: values.email,
            campaignId: campaignId
        });
        if (response.status === 201) {
            toast({
                description: 'Cliente adicionado com sucesso!',
                status: 'success',
                duration: 5000,
                isClosable: true
            });
            onUpdateState();
            onClose();

        }
        console.log(response);
    }

    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <Flex
                as="form"
                onSubmit={handleSubmit(onHandleSubmit)}
                justify="column">
                <ModalContent
                    px="4"
                    py="2"
                    bg="gray.700"
                >
                    <ModalHeader>Adicionar cliente</ModalHeader>
                    <ModalCloseButton />
                    <Stack spacing="4">
                        <Input
                            {...register('name')}
                            type="text"
                            label='Nome'
                            error={errors.name}

                        />
                        <Input
                            {...register('email')}
                            type="email"
                            label='E-mail'
                            error={errors.email}
                        />

                    </Stack>
                    <ModalFooter>
                        <Button
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="gray.600"
                            mr={3}
                            onClick={onClose}>
                            Fechar
                        </Button>
                        <Button
                            type="submit"
                            isLoading={isSubmitting}
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="blue.900">Salvar</Button>
                    </ModalFooter>
                </ModalContent>
            </Flex>
        </Modal >
    )
}