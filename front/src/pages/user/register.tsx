import React, { useState } from 'react';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';

import { Heading } from '@chakra-ui/react';
import { sendForm } from '../../api/auth-requests';
import { endpoints } from '../../api/endpoints';
import NgrForm from '../../components/ngr-form';
import useTitle from '../../hooks/title';
import { ApiResponse } from '../../types/api-response';
import { SignUpInput, UserDTO } from '../../types/generated/models';
import { Loading } from './loading';

export default function Register() {
  // Set the title of the page
  useTitle('Registrarse');
  // Recover the navigate function from the context
  const navigate = useNavigate();

  // Create a state to store the form data
  const [formData, setFormData] = useState<SignUpInput>();

  // Create a mutation to send the form data to the server
  const mutation = useMutation({
    mutationFn: (data: SignUpInput) => sendForm<SignUpInput, ApiResponse<UserDTO>>(endpoints.register, data),
    onSuccess: (data) => {
      console.log('Usuario Creado con éxito');
      console.log(data, 'data');
      navigate('/login');
    },
    onError: (error) => {
      console.log(error, 'error');
      navigate('/');
    },
  });

  // If the mutation is loading, show the loading component
  if (mutation.isLoading) return <Loading />;

  return (
    <React.Fragment>
      <Heading as="h1" size="xl" textAlign="center" my="10">
        Registrarse
      </Heading>
      <NgrForm
        schemaName="SignUpInput"
        formData={formData}
        setFormData={setFormData}
        onSubmit={(data) => {
          mutation.mutate(data);
        }}
      />
    </React.Fragment>
  );
}
