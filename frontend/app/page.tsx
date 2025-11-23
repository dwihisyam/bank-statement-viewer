import React from "react";
import FileUploader from "../components/FileUploader";
import BalanceCard from "../components/BalanceCard";
import Table from "../components/Table";
import dynamic from "next/dynamic";
import { fetchBalance } from "../utlis/app";

export default async function Page() {
  // Server component: fetch balance from backend at render time (fast on App Router)
  const balanceResp = await fetchBalance();
  const balance =
    balanceResp?.status === "success" ? balanceResp.data.balance : 0;

  return (
    <div>
      <h1 style={{ marginBottom: 12 }}>Bank Statement Viewer</h1>

      <FileUploader />

      <BalanceCard balance={balance} />

      <Table />
    </div>
  );
}
