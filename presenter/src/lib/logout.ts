import { redirect } from '@tanstack/react-router';
import Cookies from 'js-cookie';
export function Logout() {
    Cookies.remove('Authorization')
    return redirect({to: '/'})
}