import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';
import { FormControl, FormLabel, IconButton, Input, InputGroup, InputRightElement } from '@chakra-ui/react';
import { WidgetProps } from '@rjsf/utils';
import { useState } from 'react';

export default function PasswordInput(props: WidgetProps) {
  const [show, setShow] = useState(false);
  const handleClick = () => setShow(!show);
  return (
    <FormControl>
      <FormLabel>{props.label}</FormLabel>
      <InputGroup size="md">
        <Input
          id={props.id}
          pr="4.5rem"
          type={show ? 'text' : 'password'}
          onChange={(e) => props.onChange(e.target.value)}
        />
        <InputRightElement width="2.5rem">
          <IconButton
            h="1.75rem"
            size="sm"
            onClick={handleClick}
            icon={show ? <ViewOffIcon /> : <ViewIcon />}
            aria-label="Controlar la visibilidad del password"
          />
        </InputRightElement>
      </InputGroup>
    </FormControl>
  );
}
