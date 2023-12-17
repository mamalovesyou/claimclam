import { useEffect, useState } from "react";

export default function useDebounce<T>(value: T, delay: number = 500) {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  // If value is an object, there aren’t any guarantees the object ref would be the same so
  // use JSON.stringify to create a unique key
  const depValue = ["string", "number", "undefined", "null"].includes(
    typeof value
  )
    ? value
    : JSON.stringify(value);

  useEffect(() => {
    const handler: NodeJS.Timeout = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);

    // Cancel the timeout if value changes (also on delay change or unmount)
    return () => {
      clearTimeout(handler);
    };
  }, [depValue, delay]);

  return debouncedValue;
}
