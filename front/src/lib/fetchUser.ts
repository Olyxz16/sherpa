import { useUserStore } from '@/stores/userStore'

type UserData = {
  username: string,
  avatarUrl: string,
  repositories: string[]
}

export default async function fetchUser(cookie: string) {
  const userStore = useUserStore();
  const resp = await fetch("/user", {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${cookie}`
    }
  });
  const json = await resp.json() as string;
  const data = JSON.parse(json) as UserData;
  userStore.login(data.username, data.avatarUrl, data.repositories)
}
