package tk.sms

import android.app.Activity
import android.os.Bundle
import android.provider.Telephony
import android.widget.TextView


class ResultActivity : Activity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val tv: TextView = findViewById(R.id.text)
        for (pdu in Telephony.Sms.Intents.getMessagesFromIntent(intent)) {
            tv.append(pdu.displayMessageBody)
        }
    }
}