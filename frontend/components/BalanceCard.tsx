import React from "react";

export default function BalanceCard({ balance }: { balance: number }) {
  return (
    <div className="card mt-5">
      <h3>Total Balance</h3>
      <div style={{ fontSize: 22, fontWeight: 700 }}>
        Rp {balance.toLocaleString()}
      </div>
    </div>
  );
}
