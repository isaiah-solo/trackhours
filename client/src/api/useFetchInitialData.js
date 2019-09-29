import {useEffect, useState} from 'react';

import {fetchImpl} from './baseAPI';

export const useFetchInitialData = (apiPath) => {
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const fetchData = async () => {
    setIsLoading(true);
    try {
      const httpResponse = await fetchImpl(apiPath);
      const json = await httpResponse.json();
      setResponse(json);
      setIsLoading(false);
    } catch (error) {
      setError(error);
    }
  };
  useEffect(
    () => {
      fetchData();
    },
    [],
  );
  return {error, isLoading, response};
};
