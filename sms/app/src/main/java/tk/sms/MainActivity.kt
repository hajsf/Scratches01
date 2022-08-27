package tk.sms

import android.annotation.SuppressLint
import android.app.Activity
import android.app.PendingIntent
import android.content.Intent
import android.os.Bundle
import android.telephony.SmsManager
import android.widget.TextView


class MainActivity : Activity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val mgr: SmsManager = getSystemService(SmsManager::class.java) // SmsManager.getDefault()
        val token: String = mgr.createAppSpecificSmsToken(buildPendingIntent())
        val tv: TextView = findViewById(R.id.text)
        tv.text = token //getString(R.string.msg, token)
    }

    @SuppressLint("UnspecifiedImmutableFlag")
    private fun buildPendingIntent(): PendingIntent {
        return PendingIntent.getActivity(applicationContext, 1337,
                Intent(this, ResultActivity::class.java),
            PendingIntent.FLAG_IMMUTABLE or PendingIntent.FLAG_UPDATE_CURRENT) // setting the mutability flag
    }
}