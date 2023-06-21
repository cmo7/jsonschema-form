import { useEffect, useState } from 'react';

export default function <T>(key: string, defaultValue: T | null): [T, (value: T) => void] {
  const [value, setValue] = useState<T>(() => {
    try {
      const storedValue = localStorage.getItem(key);
      if (storedValue !== null) {
        return JSON.parse(storedValue);
      }
      return defaultValue;
    } catch {
      return defaultValue;
    }
  });

  useEffect(() => {
    if (value === null) {
      localStorage.removeItem(key);
      return;
    }
    const rawValue = JSON.stringify(value);
    localStorage.setItem(key, rawValue);
    localStorage.removeItem(key);
  }, [key, value]);

  return [value, setValue];
}
