import React from "react";

export default function BalanceCard({ balance }: { balance: number }) {
  return (
    <div className="card mt-5 w-3/4" style={{ justifySelf: "anchor-center" }}>
      <h3 className="text-xl font-semibold mb-2 text-gray-800">
        Total Balance
      </h3>
      <div style={{ fontSize: 22 }}>Rp {balance.toLocaleString()}</div>
    </div>
  );
}
