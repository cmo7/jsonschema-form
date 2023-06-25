import Form from '@rjsf/chakra-ui';
import validator from '@rjsf/validator-ajv8';
import { useEffect, useState } from 'react';
import { useMutation, useQuery } from 'react-query';
import { useNavigate } from 'react-router-dom';

import { sendForm } from '../../api/auth-requests';
import { endpoints } from '../../api/endpoints';
import { getSchema, getUiSchema } from '../../api/form-schema-requests';
import { useAuth } from '../../hooks/auth';
import { customWidgets } from '../../rjsf-config/widgets';
import { ApiResponse } from '../../types/api-response';
import { SignUpInput, UserResponse } from '../../types/generated/models';
import { Loading } from './loading';

export default function Register() {
  const auth = useAuth();
  const navigate = useNavigate();
  const [formData, setFormData] = useState<SignUpInput>();
  const mutation = useMutation({
    mutationFn: (data: SignUpInput) => sendForm<SignUpInput, ApiResponse<UserResponse>>(endpoints.register, data),
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

  useEffect(() => {
    if (auth.isAuthenticated()) {
      navigate('/');
    }
  }, [auth, navigate]);

  const registerSchema = useQuery('schema', async () => getSchema('SignUpInput'));
  const uiSchema = useQuery('uiSchema', async () => getUiSchema('SignUpInput'));

  if (registerSchema.isLoading || uiSchema.isLoading) return <Loading />;
  if (mutation.isLoading) return <Loading />;

  return (
    <Form
      schema={registerSchema.data}
      uiSchema={uiSchema.data}
      validator={validator}
      formData={formData}
      widgets={customWidgets}
      onSubmit={() => {
        console.log(formData, 'form data');
        mutation.mutate(formData);
      }}
      onChange={(e) => {
        setFormData(e.formData);
      }}
      liveValidate
      noHtml5Validate
    />
  );
}
