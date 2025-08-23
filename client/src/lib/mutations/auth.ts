import { signInWithGoogle, type SignInResponse } from '@/api/auth';
import { useMutation } from '@tanstack/react-query';

interface UseGoogleSignInOptions {
  onSuccess?: (data: SignInResponse) => void;
  onError?: (error: Error) => void;
}

export const useGoogleSignIn = ({ onSuccess, onError }: UseGoogleSignInOptions = {}) => {
  const {
    mutate: signIn,
    mutateAsync: signInAsync,
    isPending,
    isError,
    error,
    data,
  } = useMutation<
    SignInResponse,
    Error,
    string
  >({
    mutationFn: (idToken: string) => signInWithGoogle(idToken),
    onSuccess: (data) => {
      if (onSuccess) {
        onSuccess(data);
      }
    },
    onError: (error) => {
      if (onError) {
        onError(error);
      }
    },
  });

  return {
    signIn,
    signInAsync,
    isPending,
    isError,
    error,
    data,
  };
};
