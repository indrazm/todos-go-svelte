export async function load() {
	const res = await fetch('https://backend-go.fly.dev/');
	const data = await res.json();
	return { data };
}
