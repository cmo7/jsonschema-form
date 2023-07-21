import { Heading } from '@chakra-ui/react';
import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { sendForm } from '../../api/auth-requests';
import { endpoints } from '../../api/endpoints';
import NgrForm from '../../components/ngr-form';
import { useAuth } from '../../hooks/auth';
import useTitle from '../../hooks/title';
import { ApiResponse } from '../../types/api-response';
import { LogInInput, UserDTO } from '../../types/generated/models';
import { Loading } from './loading';

export default function Login() {
  useTitle('Iniciar sesión');
  const auth = useAuth();
  const navigate = useNavigate();
  const [formData, setFormData] = useState<LogInInput>();

  type Response = ApiResponse<{
    user: UserDTO;
    token: string;
  }>;

  const mutation = useMutation({
    mutationFn: (data: LogInInput) => sendForm<LogInInput, Response>(endpoints.login, data),
    onSuccess: (response) => {
      if (response.status === 'error') {
        console.log(response.message, 'error');
        return;
      }
      if (response.data.token) {
        auth.login(response.data.token, response.data.user);
      }
      navigate('/');
    },
    onError: (error) => {
      console.log(error, 'error');
    },
  });

  if (mutation.isLoading) return <Loading />;
  return (
    <React.Fragment>
      <Heading as="h1" size="xl" textAlign="center" my="10">
        Iniciar sesión
      </Heading>
      <NgrForm
        schemaName="LogInInput"
        formData={formData}
        setFormData={setFormData}
        onSubmit={(data) => {
          mutation.mutate(data);
        }}
      />
    </React.Fragment>
  );
}
