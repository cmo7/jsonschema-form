import {
  Button,
  CloseButton,
  Flex,
  HStack,
  IconButton,
  Link,
  Spacer,
  VStack,
  chakra,
  useColorMode,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import { AiFillHome, AiOutlineInbox, AiOutlineMenu } from 'react-icons/ai';
import { BsFillCameraVideoFill } from 'react-icons/bs';
import { FaMoon, FaSun } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';

const mainMenu = [
  {
    name: 'Dashboard',
    icon: AiFillHome,
    href: '/',
  },
  {
    name: 'Profile',
    icon: AiOutlineInbox,
    href: '/profile',
  },
  {
    name: 'Users',
    icon: BsFillCameraVideoFill,
    href: '/users',
  },
];

export default function Header() {
  const { toggleColorMode: toggleMode } = useColorMode();
  const text = useColorModeValue('dark', 'light');
  const SwitchIcon = useColorModeValue(FaMoon, FaSun);
  const bg = useColorModeValue('white', 'gray.800');
  const ref = React.useRef(null);
  const navigate = useNavigate();

  const cl = useColorModeValue('gray.800', 'white');
  const mobileNav = useDisclosure();

  const MobileNavContent = (
    <VStack
      pos="absolute"
      top={0}
      left={0}
      right={0}
      display={mobileNav.isOpen ? 'flex' : 'none'}
      flexDirection="column"
      p={2}
      pb={4}
      m={2}
      bg={bg}
      spacing={3}
      rounded="sm"
      shadow="sm"
    >
      <CloseButton aria-label="Close menu" justifySelf="self-start" onClick={mobileNav.onClose} />
      {mainMenu.map((link) => (
        <Button
          w="full"
          variant="ghost"
          leftIcon={<link.icon />}
          key={link.name}
          onClick={() => {
            navigate(link.href);
          }}
        >
          {link.name}
        </Button>
      ))}
    </VStack>
  );
  return (
    <React.Fragment>
      <chakra.header ref={ref} bg={bg} w="full" overflowY="hidden">
        <chakra.div h="4.5rem" mx="auto" maxW="container.xl">
          <Flex w="full" h="full" px="6" alignItems="center" justifyContent="space-between">
            <Flex align="flex-start">
              <Link href="/">
                <HStack mr={4}>
                  {' '}
                  <chakra.a
                    href="#"
                    fontSize="xl"
                    fontWeight="bold"
                    color="gray.600"
                    _dark={{
                      color: 'white',
                      _hover: {
                        color: 'gray.300',
                      },
                    }}
                    _hover={{
                      color: 'gray.700',
                    }}
                  >
                    Brand
                  </chakra.a>
                </HStack>
              </Link>
            </Flex>
            <Flex>
              <HStack
                spacing="5"
                display={{
                  base: 'none',
                  md: 'flex',
                }}
              >
                {mainMenu.map((link) => (
                  <Button
                    key={link.name}
                    variant="ghost"
                    size="sm"
                    _hover={{
                      color: cl,
                    }}
                    onClick={() => {
                      navigate(link.href);
                    }}
                  >
                    {link.name}
                  </Button>
                ))}
              </HStack>
            </Flex>
            <Spacer />
            <Flex justify="flex-end" align="center" color="gray.400">
              <HStack
                spacing="5"
                display={{
                  base: 'none',
                  md: 'flex',
                }}
              >
                <Button
                  colorScheme="brand"
                  variant="ghost"
                  size="sm"
                  _hover={{
                    color: cl,
                  }}
                  onClick={() => {
                    navigate('/login');
                  }}
                >
                  Acceso
                </Button>
                <Button
                  colorScheme="brand"
                  variant="ghost"
                  size="sm"
                  _hover={{
                    color: cl,
                  }}
                  onClick={() => {
                    navigate('/register');
                  }}
                >
                  Registro
                </Button>
              </HStack>
              <IconButton
                size="md"
                fontSize="lg"
                aria-label={`Switch to ${text} mode`}
                variant="ghost"
                color="current"
                ml={{
                  base: '0',
                  md: '3',
                }}
                onClick={toggleMode}
                icon={<SwitchIcon />}
              />
              <IconButton
                display={{
                  base: 'flex',
                  md: 'none',
                }}
                aria-label="Open menu"
                fontSize="20px"
                color="gray.800"
                _dark={{
                  color: 'inherit',
                }}
                variant="ghost"
                icon={<AiOutlineMenu />}
                onClick={mobileNav.onOpen}
              />
            </Flex>
          </Flex>
          {MobileNavContent}
        </chakra.div>
      </chakra.header>
    </React.Fragment>
  );
}
