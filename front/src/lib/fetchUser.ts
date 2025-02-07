import { useUserStore } from '@/stores/userStore'
import type { UserData } from '@/types/UserData'

export default async function fetchUser(cookie: string) : Promise<boolean> {
  const userStore = useUserStore();
  try {
    const resp = await fetch("/user", {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${cookie}`
      }
    });
    const data : UserData = await resp.json() as UserData;
    userStore.login(data.username, data.avatarUrl, data.repositories)
  } catch(e) {
    return false;
  }
  return true;
}
