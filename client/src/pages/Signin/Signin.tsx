import { useGoogleSignIn } from '@/lib/mutations/auth';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { GoogleLogin } from '@react-oauth/google';
import { toast } from 'sonner';
import { useNavigate } from '@tanstack/react-router';

interface CredentialResponse {
  credential?: string;
}

export default function SigninPage() {
  const navigate = useNavigate();
  
  const { signIn, isPending } = useGoogleSignIn({
    onSuccess: (data) => {
      toast.success(`Welcome ${data.user.full_name}`);
      localStorage.setItem('accessToken', data.access_token);
      navigate({ to: "/" });
    },
    onError: () => {
      toast.error("Sign-in failed.");
    }
  });

  const handleGoogleSuccess = (credentialResponse: CredentialResponse) => {
    if (credentialResponse.credential) {
      signIn(credentialResponse.credential);
    } else {
      toast.error("Sign-in failed.");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50">
      <Card className="w-[350px] shadow-lg">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl font-bold text-center">Welcome</CardTitle>
          <CardDescription className="text-center">
            Sign in to access your account
          </CardDescription>
        </CardHeader>
        <CardContent className="flex flex-col items-center justify-center">
          {isPending ? (
            <div className="flex items-center justify-center py-3">
              <svg className="animate-spin h-8 w-8 text-primary" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span className="ml-3">Signing in...</span>
            </div>
          ) : (
            <GoogleLogin
              onSuccess={handleGoogleSuccess}
              onError={() => toast.error("Google login failed.")}
              width="300px"
            />
          )}
        </CardContent>
        <CardFooter className="flex flex-col items-center justify-center gap-2 text-xs text-gray-500">
          <p>By signing in, you agree to our</p>
          <div className="flex gap-1">
            <a href="https://github.com/adistrim/usDrive/blob/main/ToS" className="text-blue-500 hover:underline">Terms of Service</a>
            <span>and</span>
            <a href="https://github.com/adistrim/usDrive/blob/main/PRIVACY" className="text-blue-500 hover:underline">Privacy Policy</a>
          </div>
        </CardFooter>
      </Card>
    </div>
  );
}
