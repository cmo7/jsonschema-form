import type { Meta, StoryObj } from '@storybook/react';

import UserBadge, { UserBadgeProps } from '../components/user/user-badge';

const meta: Meta<typeof UserBadge> = {
  title: 'Example/UserBadge',
  component: UserBadge,
  tags: ['autodocs'],
};

export default meta;

type Story = StoryObj<typeof UserBadge>;

export const Default: Story = {
  args: {
    id: 1,
    email: 'mars@rider.com',
    friends: [],
  },
};
