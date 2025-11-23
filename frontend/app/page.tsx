"use client";

import React, { useEffect, useState } from "react";
import FileUploader from "../components/FileUploader";
import BalanceCard from "../components/BalanceCard";
import Table from "../components/Table";
import { fetchBalance } from "../lib/api";

export const dynamic = "force-dynamic";

export default function Page() {
  const [balance, setBalance] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function getBalance() {
      try {
        const resp = await fetchBalance();
        const bal = resp?.status === "success" ? resp.data.balance : 0;
        setBalance(bal);
      } catch (err) {
        console.error("Failed to fetch balance:", err);
        setBalance(0);
      } finally {
        setLoading(false);
      }
    }

    getBalance();
  }, []);

  const hasData = balance !== null && balance > 0;

  return (
    <div>
      <h1 style={{ fontSize: 22, fontWeight: 700, textAlign: "center" }}>
        Bank Statement Viewerr
      </h1>

      <FileUploader />

      {hasData && (
        <>
          <BalanceCard balance={balance} />
          <Table />
        </>
      )}
    </div>
  );
}
