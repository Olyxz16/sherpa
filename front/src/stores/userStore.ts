import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useUserStore = defineStore('userStore', () => {
  const loggedIn = ref(false);
  const username = ref('');
  const avatarUrl = ref('');
  const repositories = ref<string[]>([]);

  function login(usernameParam: string, avatarUrlParam: string, repositoriesParam: string[]) {
    loggedIn.value = true;
    username.value = usernameParam;
    avatarUrl.value = avatarUrlParam;
    repositories.value = repositoriesParam;
  }
  return { loggedIn, username, avatarUrl, repositories, login };
});
