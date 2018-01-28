package kermit;

import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.TextField;
import javafx.scene.text.Text;
import javafx.stage.Stage;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class SendController {

  private boolean buttonArmed;
  private Stage stage;

  @FXML
  private TextField paymentField;

  @FXML
  private TextField payeeAddress;

  @FXML
  private Text availableTracker;

  @FXML
  private Button sendButton;

  public void init(float balance, Stage stage) {
    deactivateButton();
    this.stage = stage;
    availableTracker.setText(String.format(java.util.Locale.UK,"%.2f", balance));
    paymentField.textProperty().addListener((observable, oldValue, newValue) -> {
      float amountToPay;
      if (!newValue.isEmpty()) {
        if (!newValue.matches("\\d*\\.?\\d*")) {
          paymentField.setText(oldValue);
        } else {
          if (Float.parseFloat(newValue) > balance) {
            paymentField.setText(String.valueOf(balance));
          }
        }
        amountToPay = Float.parseFloat(paymentField.getText());
      } else {
        amountToPay = 0;
      }
      float available = balance - amountToPay;
      availableTracker.setText(String.format(java.util.Locale.UK,"%.2f", available));

      updateButton();

    });

    payeeAddress.textProperty().addListener((observable, oldValue, newValue) -> {
      if (!Pattern.matches("[0-9a-fA-FxX]{0,18}", newValue)) {
        payeeAddress.setText(oldValue);
      }
      updateButton();
    });

  }

  private void updateButton() {
    Pattern p = Pattern.compile("^0[Xx][0-9a-fA-F]{16}$");
    Matcher m = p.matcher(payeeAddress.getText());
    boolean buttonActive = !paymentField.getText().isEmpty() &&
                           m.matches();
    if (buttonActive) {
      activateButton();
    } else {
      deactivateButton();
    }
  }

  private void activateButton() {
    buttonArmed = true;
    sendButton.setStyle("-fx-border-color: rgb(253, 255, 253); -fx-text-fill: rgb(253, 255, 253);");

  }

  private void deactivateButton() {
    buttonArmed = false;
    sendButton.setStyle("-fx-border-color: rgb(79, 129, 59); -fx-text-fill: rgb(79, 129, 59);");

  }

  @FXML
  private void handleSendButtonAction() {
    if (buttonArmed) {
      System.out.println("boop");
      stage.close();
    }
  }

}
