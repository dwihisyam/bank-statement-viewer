"use client";
import React, { useEffect, useState } from "react";
import { fetchIssues } from "../lib/api";

type Tx = {
  timestamp: number;
  name: string;
  type: string;
  amount: number;
  status: string;
  description: string;
};

export default function IssuesTableClient() {
  const [issues, setIssues] = useState<Tx[]>([]);
  const [page, setPage] = useState(1);
  const [pageSize] = useState(10);
  const [sortField, setSortField] = useState<keyof Tx | null>(null);
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("asc");

  useEffect(() => {
    (async () => {
      const res = await fetchIssues();
      if (res.status === "success") setIssues(res.data || []);
    })();
  }, []);

  // handle click header to sort
  const handleSort = (field: keyof Tx) => {
    if (sortField === field) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortField(field);
      setSortOrder("asc");
    }
    setPage(1); // reset to first page when sort changes
  };

  // sorting
  const sortedIssues = [...issues].sort((a, b) => {
    if (!sortField) return 0;
    let valA = a[sortField];
    let valB = b[sortField];

    if (typeof valA === "string") valA = valA.toLowerCase();
    if (typeof valB === "string") valB = valB.toLowerCase();

    if (valA < valB) return sortOrder === "asc" ? -1 : 1;
    if (valA > valB) return sortOrder === "asc" ? 1 : -1;
    return 0;
  });

  // pagination
  const total = issues.length;
  const pages = Math.max(1, Math.ceil(total / pageSize));
  const start = (page - 1) * pageSize;
  const paged = sortedIssues.slice(start, start + pageSize);

  const getSortIndicator = (field: keyof Tx) => {
    if (sortField !== field) return "";
    return sortOrder === "asc" ? "↑" : "↓";
  };

  return (
    <div className="card w-3/4" style={{ justifySelf: "anchor-center" }}>
      <h3 className="text-xl font-semibold mb-4 text-gray-800">
        Non-Success Transactions (Issues)
      </h3>
      <table className="table" style={{ marginTop: 8 }}>
        <thead>
          <tr>
            <th onClick={() => handleSort("timestamp")}>
              Time {getSortIndicator("timestamp")}
            </th>
            <th onClick={() => handleSort("name")}>
              Name {getSortIndicator("name")}
            </th>
            <th onClick={() => handleSort("type")}>
              Type {getSortIndicator("type")}
            </th>
            <th onClick={() => handleSort("amount")}>
              Amount {getSortIndicator("amount")}
            </th>
            <th onClick={() => handleSort("status")}>
              Status {getSortIndicator("status")}
            </th>
            <th onClick={() => handleSort("description")}>
              Description {getSortIndicator("description")}
            </th>
          </tr>
        </thead>
        <tbody>
          {paged.map((tx, i) => (
            <tr key={i}>
              <td>{new Date(tx.timestamp * 1000).toLocaleString()}</td>
              <td>{tx.name}</td>
              <td>{tx.type}</td>
              <td>Rp. {tx.amount.toLocaleString()}</td>
              <td className="text-center">
                <span
                  className={`inline-flex items-center gap-1 px-2 py-1 rounded-lg text-xs font-semibold text-center align-center 
      ${
        tx.status === "FAILED"
          ? "bg-red-100 text-red-700"
          : "bg-yellow-100 text-yellow-700"
      }`}
                >
                  {tx.status === "FAILED" ? "❌" : "⚠️"}
                  {tx.status}
                </span>
              </td>
              <td>{tx.description}</td>
            </tr>
          ))}
          {paged.length === 0 && (
            <tr>
              <td colSpan={6} style={{ textAlign: "center", color: "#6b7280" }}>
                No issues
              </td>
            </tr>
          )}
        </tbody>
      </table>

      <div
        style={{
          display: "flex",
          gap: 8,
          marginTop: 17,
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <button
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          disabled={page === 1}
          className="px-3 py-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
          title="Previous page"
        >
          ←
        </button>

        <div style={{ color: "#374151" }}>
          Page {page} / {pages}
        </div>

        <button
          onClick={() => setPage((p) => Math.min(pages, p + 1))}
          disabled={page === pages}
          className="px-3 py-1 rounded bg-gray-200 hover:bg-gray-300 disabled:opacity-50"
          title="Next page"
        >
          →
        </button>
      </div>
    </div>
  );
}
