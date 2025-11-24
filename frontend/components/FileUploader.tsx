"use client";
import React, { useState } from "react";
import { uploadCSV } from "../utils/api";
import Snackbar from "./Snackbar";

export default function FileUploader() {
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  const [snackbar, setSnackbar] = useState<{
    message: string;
    type: "success" | "error";
  } | null>(null);

  function showSnackbar(message: string, type: "success" | "error") {
    setSnackbar({ message, type });
  }

  async function onUpload() {
    if (!file) {
      showSnackbar("Please select a CSV file first", "error");
      return;
    }
    setLoading(true);

    try {
      const resp = await uploadCSV(file);
      if (resp.status === "success") {
        showSnackbar(`Upload successful: ${resp.data.count} rows`, "success");
        setTimeout(() => window.location.reload(), 800);
      } else {
        showSnackbar(resp.message || "Upload failed", "error");
      }
    } catch (e) {
      showSnackbar("Upload error: " + String(e), "error");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="max-w-md mx-auto mt-5 bg-white shadow-lg rounded-2xl p-6 border border-gray-100">
      <h3 className="text-xl font-semibold mb-4 text-gray-800">Upload CSV</h3>

      <label className="relative flex flex-col items-center justify-center w-full h-40 border-2 border-dashed border-gray-300 rounded-xl cursor-pointer hover:border-blue-500 transition">
        <div className="flex flex-col items-center gap-2">
          <span className="text-gray-500 text-sm">
            {file ? file.name : "Click to choose CSV file"}
          </span>
          <span className="text-blue-600 text-sm font-medium">
            {file ? "File ready to upload" : "Browse from computer"}
          </span>
        </div>
        <input
          type="file"
          className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
          accept=".csv,text/csv"
          onChange={(e) => setFile(e.target.files?.[0] ?? null)}
        />
      </label>

      <button
        onClick={onUpload}
        disabled={loading}
        className={`w-full mt-5 py-3 rounded-xl font-semibold text-white transition
          ${
            loading
              ? "bg-blue-300 cursor-not-allowed"
              : "bg-blue-600 hover:bg-blue-700 active:scale-95"
          }
        `}
      >
        {loading ? "Uploading..." : "Upload"}
      </button>

      {snackbar && (
        <Snackbar
          message={snackbar.message}
          type={snackbar.type}
          onClose={() => setSnackbar(null)}
        />
      )}
    </div>
  );
}
