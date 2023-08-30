export async function load() {
	const res = await fetch('http://localhost:8080/api/v1/task');
	const data = await res.json();
	return { data };
}
