export const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://backend:8080";

export async function fetchBalance() {
  try {
    const res = await fetch(`${API_URL}/balance`, { cache: "no-store" });
    return await res.json();
  } catch (err) {
    return { status: "error", message: String(err) };
  }
}

export async function fetchIssues() {
  try {
    const res = await fetch(`${API_URL}/issues`, { cache: "no-store" });
    return await res.json();
  } catch (err) {
    return { status: "error", message: String(err), data: [] };
  }
}

export async function uploadCSV(file: File) {
  const form = new FormData();
  form.append("file", file);
  try {
    const res = await fetch(`${API_URL}/upload`, {
      method: "POST",
      body: form,
    });
    return res.json();
  } catch (err) {
    return { status: "error", message: String(err) };
  }
}
