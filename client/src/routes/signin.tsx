import SigninPage from '@/pages/Signin/Signin'
import { createFileRoute } from '@tanstack/react-router'
import { GoogleOAuthProvider } from '@react-oauth/google'
import { G_CLIENT } from '@/config'

export const Route = createFileRoute('/signin')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <GoogleOAuthProvider clientId={G_CLIENT}>
      <SigninPage />
    </GoogleOAuthProvider>
  )
}