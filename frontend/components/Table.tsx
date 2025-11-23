"use client";
import React, { useEffect, useState } from "react";
import { fetchIssues } from "../utlis/app";

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

  useEffect(() => {
    (async () => {
      const res = await fetchIssues();
      if (res.status === "success") setIssues(res.data || []);
    })();
  }, []);

  // simple client-side pagination
  const total = issues.length;
  const pages = Math.max(1, Math.ceil(total / pageSize));
  const start = (page - 1) * pageSize;
  const paged = issues.slice(start, start + pageSize);

  return (
    <div className="card">
      <h3>Non-Success Transactions (Issues)</h3>
      <table className="table" style={{ marginTop: 8 }}>
        <thead>
          <tr>
            <th>Time</th>
            <th>Name</th>
            <th>Type</th>
            <th>Amount</th>
            <th>Status</th>
            <th>Description</th>
          </tr>
        </thead>
        <tbody>
          {paged.map((tx, i) => (
            <tr key={i}>
              <td>{new Date(tx.timestamp * 1000).toLocaleString()}</td>
              <td>{tx.name}</td>
              <td>{tx.type}</td>
              <td>{tx.amount.toLocaleString()}</td>
              <td
                className={
                  tx.status === "FAILED" ? "status-failed" : "status-pending"
                }
              >
                {tx.status}
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
        style={{ display: "flex", gap: 8, marginTop: 12, alignItems: "center" }}
      >
        <button
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          disabled={page === 1}
        >
          Prev
        </button>
        <div style={{ color: "#374151" }}>
          Page {page} / {pages}
        </div>
        <button
          onClick={() => setPage((p) => Math.min(pages, p + 1))}
          disabled={page === pages}
        >
          Next
        </button>
      </div>
    </div>
  );
}
