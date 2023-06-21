import Form from '@rjsf/chakra-ui';
import { UiSchema } from '@rjsf/utils';
import validator from '@rjsf/validator-ajv8';
import { useEffect, useState } from 'react';
import { useMutation, useQuery } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { getSchema, sendForm } from '../api/auth-requests';
import { endpoints } from '../api/endpoints';
import { useAuth } from '../hooks/auth';
import { customWidgets } from '../rjsf-config/widgets';
import { ApiResponse } from '../types/api-response';
import { LogInInput, UserResponse } from '../types/generated/models';
import { Loading } from './loading';

const uiSchema: UiSchema = {
  email: {
    'ui:widget': 'email',
  },
  password: {
    'ui:widget': 'password',
  },
};

export default function Login() {
  const auth = useAuth();
  const navigate = useNavigate();
  const [formData, setFormData] = useState<LogInInput>();

  type Response = ApiResponse<{
    user: UserResponse;
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

  useEffect(() => {
    if (auth.isAuthenticated()) {
      navigate('/');
    }
  }, [auth, navigate]);

  const loginSchema = useQuery('schema', async () => getSchema(endpoints.user, 'LogInInput'));

  if (loginSchema.isLoading) return <Loading />;
  if (loginSchema.isError) return <div>error</div>;
  if (mutation.isLoading) return <Loading />;
  return (
    <Form
      schema={loginSchema.data}
      uiSchema={uiSchema}
      validator={validator}
      widgets={customWidgets}
      formData={formData}
      onSubmit={() => {
        mutation.mutate(formData);
      }}
      onChange={(e) => {
        setFormData(e.formData);
      }}
    />
  );
}
