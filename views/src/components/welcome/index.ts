import router from '@/router'

export default async function submitMasterkey(masterkey: string) {
  console.log(masterkey);
  const data = {
    'masterkey': masterkey
  }
  const resp = await fetch('/auth/masterkey', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });
  router.push({path: '/'})
}
