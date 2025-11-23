"use client";
import React, { useState } from "react";
import { uploadCSV, fetchBalance } from "../utlis/app";

export default function FileUploader() {
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  async function onUpload() {
    if (!file) return alert("Pilih file CSV terlebih dahulu");
    setLoading(true);
    try {
      const resp = await uploadCSV(file);
      if (resp.status === "success") {
        alert(`Upload sukses: ${resp.data.count} rows`);
        // trigger client-side refresh for balance/issue table by reloading window
        // (simplest approach). Alternatively use mutate/store.
        window.location.reload();
      } else {
        alert("Upload gagal: " + (resp.message || "unknown"));
      }
    } catch (e) {
      alert("Upload error: " + String(e));
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="card">
      <h3>Upload CSV</h3>
      <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
        <input
          type="file"
          accept=".csv,text/csv"
          onChange={(e) => setFile(e.target.files?.[0] ?? null)}
        />
        <button onClick={onUpload} disabled={loading}>
          {loading ? "Uploading..." : "Upload"}
        </button>
      </div>
      <p style={{ color: "#6b7280", marginTop: 8 }}>
        Field name form: <code>file</code>
      </p>
    </div>
  );
}
