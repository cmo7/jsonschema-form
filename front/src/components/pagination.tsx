import { ArrowLeftIcon, ArrowRightIcon } from '@chakra-ui/icons';
import { Flex, Icon, chakra, useColorModeValue } from '@chakra-ui/react';
import { PropsWithChildren, useState } from 'react';
import { HiDotsHorizontal } from 'react-icons/hi';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';

type PagButtonProps = PropsWithChildren<{
  active?: boolean;
  disabled?: boolean;
  onClick?: () => void;
}>;

function PagButton({ active, disabled, children, onClick }: PagButtonProps) {
  const activeStyle = {
    bg: 'brand.600',
    _dark: {
      bg: 'brand.500',
    },
    color: 'white',
  };
  return (
    <chakra.button
      mx={1}
      px={4}
      py={2}
      rounded="md"
      bg="white"
      _dark={{
        bg: 'gray.800',
      }}
      color="gray.700"
      opacity={disabled && 0.6}
      _hover={!disabled && activeStyle}
      cursor={disabled && 'not-allowed'}
      {...(active && activeStyle)}
      onClick={onClick}
    >
      {children}
    </chakra.button>
  );
}

type MButtonProps = {
  left?: boolean;
  right?: boolean;
};

function MButton({ left, right }: MButtonProps) {
  const DoubleArrow = left ? ArrowLeftIcon : ArrowRightIcon;
  const [hovered, setHovered] = useState(false);
  const color1 = useColorModeValue('brand.800', 'brand.700');
  const color2 = useColorModeValue('gray.100', 'gray.200');
  return (
    <chakra.a
      w={8}
      py={2}
      color="gray.700"
      _dark={{
        color: 'gray.200',
      }}
      onMouseOver={() => setHovered(true)}
      onMouseOut={() => setHovered(false)}
      cursor="pointer"
      textAlign="center"
    >
      {hovered ? (
        <Icon as={DoubleArrow} boxSize={3} cursor="pointer" color={color1} />
      ) : (
        <Icon as={HiDotsHorizontal} color={color2} boxSize={4} opacity={0.5} />
      )}
    </chakra.a>
  );
}

type PaginationProps = {
  page: number;
  pageSize: number;
  totalPages: number;
  setPage: (page: number) => void;
  setPageSize: (pageSize: number) => void;
};

export default function Pagination({ page, pageSize, totalPages, setPage, setPageSize }: PaginationProps) {
  const pages = Array.from({ length: totalPages }, (_, i) => i + 1);

  return (
    <Flex
      bg="#edf3f8"
      _dark={{
        bg: '#3e3e3e',
      }}
      //p={50}
      w="full"
      alignItems="center"
      justifyContent="center"
    >
      <Flex>
        <PagButton active={page >= 1}>
          <Icon
            as={IoIosArrowBack}
            color="gray.700"
            _dark={{
              color: 'gray.200',
            }}
            boxSize={4}
            onClick={() => setPage(page - 1)}
          />
        </PagButton>
        {pages.map((p) => (
          <PagButton key={p} active={p === page} onClick={() => setPage(p)}>
            {p}
          </PagButton>
        ))}
        <PagButton active={page <= totalPages}>
          <Icon
            as={IoIosArrowForward}
            color="gray.700"
            _dark={{
              color: 'gray.200',
            }}
            boxSize={4}
            onClick={() => setPage(page + 1)}
          />
        </PagButton>
      </Flex>
    </Flex>
  );
}
